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
	"os/user"
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
func ClonePackageRepo() {
	p := new(PackageManager)

	// Subject to change
	packages_url := "https://github.com/gopac/packages.git"

	if (checkForPackageRepo(p) != true) {
		fmt.Println(p.home_dir)
		git.Clone(packages_url, p.home_dir, new(git.CloneOptions))
	}
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
 */
func makeHomeDirectory(home_dir string) {
	if err := os.Mkdir(home_dir, 0776); err != nil {
		fmt.Println("error:", err)
	}
}
