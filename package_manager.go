/**
 * gopac v0.1-dev
 *
 * (c) Ground Six
 *
 * @package gopac
 * @version 0.1-dev
 * 
 * @author Harry Lawrence <http://github.com/hazbo>
 *
 * License: MIT
 * 
 * For the full copyright and license information, please view the LICENSE
 * file that was distributed with this source code.
 */

package main

import (
    "log"
    "fmt"
    "os"
    "sync"
    "os/user"
    "./vendor/pb"
    "./vendor/jconfig"
    git "./vendor/git2go"
)

/**
 * @var string user's home directory
 */
type PackageManager struct {
    home_dir string
}

/**
 * Clones the Gopac 'packages' directory
 * to local machine
 *
 * @return nil
 */
func ClonePackageRepo() *PackageManager {
    p := new(PackageManager)

    // Subject to change
    packages_url := "https://github.com/gopac/packages.git"

    if (checkForPackageRepo(p) != true) {
        git.Clone(packages_url, p.home_dir, new(git.CloneOptions))
    }

    return p
}

/**
 * Clones all first iteration of dependencies
 * at the moment
 * TODO: Abstract into different functions
 *
 * @param *Packages package object
 *
 * @return nil
 */
func (p *PackageManager) CloneDependencies(packages *Packages) {
    package_count := len(packages.package_map)
    bar := pb.StartNew(package_count * 2)

    fmt.Println("Installing:")
    for packet, branch := range packages.package_map {
        fmt.Println(packet, " -> ", branch)
    }

    for packet, branch := range packages.package_map {
        file_query := p.home_dir + "/" + packet + "/package.json"
        if _, err := os.Stat(file_query); err == nil {
            config_data := jconfig.LoadConfig(file_query)
            new_package := config_data.GetStringMap("package")
            makeVendorDirectory()

            clone_options := new(git.CloneOptions)
            clone_options.CheckoutBranch = branch
            packet_path := "vendor/" + new_package["vendor"].(string) + "/" + packet

            var wg sync.WaitGroup
            wg.Add(1)

            go func () {
                bar.Increment()
                git.Clone(new_package["url"].(string), packet_path, clone_options)
                bar.Increment()
                wg.Done()
            }()

            wg.Wait()
        }
    }
    bar.FinishPrint("\nSuccess!\n")
}

/**
 * Checks to see if the packages repo
 * already exists on the local machine
 *
 * @param *PackageManager
 *
 * @return bool exists or not
 */
func checkForPackageRepo(p *PackageManager) bool {
    usr, err := user.Current()
    if err != nil {
        log.Fatal( err )
    }
    p.home_dir = usr.HomeDir + "/.gopac"
    _, err = os.Stat(p.home_dir)
    if err == nil {
        return true
    }
    makeHomeDirectory(p.home_dir)
    return false
}

/**
 * Creates the Gopac home directory
 * on the user's local machine
 *
 * @param string user's home directory
 *
 * @return nil
 */
func makeHomeDirectory(home_dir string) {
    if err := os.Mkdir(home_dir, 0776); err != nil {
        fmt.Println("error:", err)
    }
}

/**
 * Creates the vendor home directory
 * on the user's local machine
 *
 * @return nil
 */
func makeVendorDirectory() {
    if _, err := os.Stat("./vendor"); err != nil {
        if err := os.Mkdir("./vendor", 0776); err != nil {
            fmt.Println("error:", err)
        }
    }
}
