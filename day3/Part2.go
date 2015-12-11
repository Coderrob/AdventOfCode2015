package main

// North 	^
// South 	v
// East 	> 
// West 	<

/*
--- Part Two ---

The next year, to speed up the process, Santa creates a robot version of himself, Robo-Santa, to deliver presents with him.

Santa and Robo-Santa start at the same location (delivering two presents to the same starting house), then take turns moving based on instructions from the elf, who is eggnoggedly reading from the same script as the previous year.

This year, how many houses receive at least one present?

For example:

^v delivers presents to 3 houses, because Santa goes north, and then Robo-Santa goes south.
^>v< now delivers presents to 3 houses, and Santa and Robo-Santa end up back where they started.
^v^v^v^v^v now delivers presents to 11 houses, with Santa going one direction and Robo-Santa going the other.
*/

import (
	"fmt"
)

func getHouseCountVistedMoreThanOnceBySantaAndRobotSanta(path string) int {
	var houseCount = 1

	var santaPositions = map[string]int { "0,0": 1 }
	var robotPositions = map[string]int { "0,0": 1 }

	var xRobotPosition int
	var yRobotPosition int 
	
	var xSantaPosition int
	var ySantaPosition int
	
	for index := 0; index < len(path); index++ {
		if((index % 2) == 0) {
			switch(path[index]) {
				case '<':
					xSantaPosition--
				case '>':
					xSantaPosition++
				case '^':
					ySantaPosition++
				case 'v':
					ySantaPosition--
			}	
			
			var coordinate = fmt.Sprintf("%d,%d", xSantaPosition, ySantaPosition)
			
			if (santaPositions[coordinate] != 0) {
				santaPositions[coordinate]++
			} else {
				santaPositions[coordinate] = 1
				if (robotPositions[coordinate] == 0) {
					houseCount++
				}
			}
		} else {
			switch(path[index]) {
				case '<':
					xRobotPosition--
				case '>':
					xRobotPosition++
				case '^':
					yRobotPosition++
				case 'v':
					yRobotPosition--
			}	
			
			var coordinate = fmt.Sprintf("%d,%d", xRobotPosition, yRobotPosition)
			
			if (robotPositions[coordinate] != 0) {
				robotPositions[coordinate]++
			} else {				
				robotPositions[coordinate] = 1
				if (santaPositions[coordinate] == 0) {
					houseCount++
				}
			}
		}
	}

	return houseCount
}

func getThreeHousesVisitedTestPath() string {
	return "^v"
}

func getThreeHousesVisitedBySantaAndRobotTestPath() string {
	return "^>v<"
}

func getElevenHousesVisistedTestPath() string {
	return "^v^v^v^v^v"
}