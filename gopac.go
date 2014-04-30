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
	"os"
	"fmt"
)

func main() {
	//config := ParseConfig()
	HandleArguments(os.Args)
}

/**
 * For testing purposes only
 */
func TestCallback() {
	fmt.Println("At callback")
}