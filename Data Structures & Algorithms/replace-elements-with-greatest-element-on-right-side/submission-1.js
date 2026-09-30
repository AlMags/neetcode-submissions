class Solution {
    /**
     * @param {number[]} arr
     * @return {number[]}
     */
    replaceElements(arr) {
        let res = new Array(arr.length).fill(0)

        for(let i = 0; i < arr.length; i++) {
            let sliceArr = arr.slice(i+1, arr.length) 
            let max = arr[i+1]
            for(let j = 0; j < arr.length; j++) {
                if(max < sliceArr[j]) {
                    max = sliceArr[j]
                }
            }
            res[i] = max
        }

        res[arr.length - 1] = -1

        return res
    }
}
