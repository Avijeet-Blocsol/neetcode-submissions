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
     * @returns {boolean}
     */
    canAttendMeetings(intervals) {
        
        if (intervals.length < 2) {
            return true
        }
        
        intervals = mergeSort(intervals)

        console.log("intervals are: ", intervals)
        let curr = intervals[0]

        for (let i = 1; i < intervals.length; i++) {
            if (intervals[i].start < curr.end) {
                return false
            } else {
                curr = intervals[i]
            }
        }

        return true
    }
}

function mergeSort(list) {
    if (list.length === 0) {
        return [];
    }

    if (list.length === 1) {
        return list;
    }

    let mid = Math.floor(list.length / 2);
    return mergeLists(mergeSort(list.slice(0, mid)), mergeSort(list.slice(mid)));
}

function mergeLists(a, b) {
    let i = 0;
    let j = 0;

    let list = [];

    while (i < a.length && j < b.length) {
        if (a[i].start < b[j].start) {
            list.push(a[i]);
            i++;
        } else {
            list.push(b[j]);
            j++;
        }
    }

    list = list.concat(a.slice(i));
    list = list.concat(b.slice(j));

    return list;
}

