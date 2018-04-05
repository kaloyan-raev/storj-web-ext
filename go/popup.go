// Copyright (C) 2018 Kaloyan Raev
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
package main

import (
	"fmt"

	"github.com/Storj/go-libstorj/storj"
	"honnef.co/go/js/dom"
)

var (
	doc = dom.GetWindow().Document()
)

func renderStatus(statusText string) {
	doc.GetElementByID("status").SetInnerHTML(statusText)
}

func main() {
	renderStatus("Loading bridge API info...")

	info, err := storj.GetInfo(storj.NewEnv())
	if err != nil {
		renderStatus(fmt.Sprintf("%v", err))
	} else {
		renderStatus(fmt.Sprintf("Title: %s<br>Description: %s<br>Version: %s<br>Host: %s\n",
			info.Title, info.Description, info.Version, info.Host))
	}
}
