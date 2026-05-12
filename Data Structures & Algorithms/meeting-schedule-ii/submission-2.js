/**
 * Definition of Interval:
 * class Interval {
 *   constructor(start, end) {
 *     this.start = start;
 *     this.end = end;
 *   }
 * }
 */

class Solution {
    /**
     * @param {Interval[]} intervals
     * @returns {number}
     */
    minMeetingRooms(intervals) {

        if (intervals.length == 0) {
            return 0
        }

        let startTimes = []
        let endTimes = []

        for (let i = 0; i < intervals.length; i++) {
            startTimes.push(intervals[i].start)
            endTimes.push(intervals[i].end)
        }

        startTimes = startTimes.sort((a, b) => {
            return a - b
        })

        endTimes = endTimes.sort((a, b) => {
            return a - b
        })

        let requiredRooms = 0
        let maxRooms = 0
        let i = 0; let j = 0

        while (i < startTimes.length) {
            if (startTimes[i] < endTimes[j]) {
                requiredRooms += 1
                maxRooms = Math.max(maxRooms, requiredRooms)
                i +=1 
            } else {
                requiredRooms -=1
                j += 1
            }
        }

        return maxRooms
    }
}
