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
    "./vendor/jconfig"
    "path/filepath"
    "log"
    "os"
    "fmt"
)

/**
 * @var string full file path for config file
 * @var map[string]string all require data
 */
type Packages struct {
    file_path   string
    package_map map[string]string
}

/**
 * Starts off the initial configuration
 * parsing process to return all require
 * values.
 *
 * @return Packages completed struct
 */
func ParseConfig() Packages {
    p := Packages{}
    if (checkGopacfileExists(&p) == true) {
        config_data := jconfig.LoadConfig(p.file_path)
        require := config_data.GetStringMap("require")

        sortRequiredPackages(&p, require)
    }
    return p
}

/**
 * Gets the path of the current directory
 * when gopac is ran.
 *
 * @string directory path
 */
func getRelativePath() string {
    dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
        log.Fatal(err)
    }
    return dir
}

/**
 * Checks to see if the gopacfile actually
 * exists first.
 *
 * @param *Packages the package struct
 */
func checkGopacfileExists(packages_data *Packages) bool {
    file_path := getRelativePath()
    file_path = file_path + "/gopac.json"
    if _, err := os.Stat(file_path); err == nil {
        packages_data.file_path = file_path
        return true
    }
    fmt.Printf("error: gopac.json not found.")

    return false
}

/**
 * Converts type interface{} for map to
 * map[string]string
 *
 * @param *Packages the package struct
 * @param map[string]interface{} the initial require data
 */
func sortRequiredPackages(p *Packages, require map[string]interface{})  {
    p.package_map = make(map[string]string)
    for k, v := range require {
        p.package_map[k] = v.(string)
    }
}
