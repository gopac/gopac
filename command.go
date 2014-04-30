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
    "fmt"
)

/**
 * Handles the os arguments taken from
 * the initial command
 *
 * @param []string all command arguments
 *
 * @return bool success or fail
 */
func HandleArguments(command []string) bool {
    if (len(command) == 2) {
        switch command[1] {
            case "install" : TestCallback()
            case "update"  : TestCallback()
        }
        return true
    }
    fmt.Printf("error: expecting 1 argument, recieved %d.\n", len(command) - 1)
    return false
}