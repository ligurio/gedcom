/*
 * timenet
 *   - http://vis.stanford.edu/papers/timenets
 *   - http://vis.berkeley.edu/courses/cs294-10-sp10/wiki/images/f/f2/Family_Tree_Visualization_-_Final_Paper.pdf
 * timeline
 *   - https://github.com/davorg/svg-timeline-genealogy
 * graphviz:
 *   - https://github.com/adrienverge/familytreemaker
 *   - https://github.com/vmiklos/ged2dot
 *   - моя реализация
 * [GEPS 030: New Visualization Techniques](https://www.gramps-project.org/wiki/index.php/GEPS_030:_New_Visualization_Techniques)
 * [Geneaquilts](https://aviz.fr/geneaquilts/)
 * https://github.com/nicolaskruchten/genealogy
 */

package main

import (
	"github.com/ajstarks/svgo"
	"os"
)

func drawTimenet() {
	width := 500
	height := 500
	canvas := svg.New(os.Stdout)
	canvas.Start(width, height)
	canvas.Circle(width/2, height/2, 100)
	canvas.Text(width/2, height/2, "Hello, SVG", "text-anchor:middle;font-size:30px;fill:white")
	canvas.End()
}
