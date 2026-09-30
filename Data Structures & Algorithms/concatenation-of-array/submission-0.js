class Solution {
    /**
     * @param {number[]} nums
     * @return {number[]}
     */
    getConcatenation(nums) {
        const n = nums.length * 2
        const res = new Array(n)

        for(let i = 0; i < nums.length; i++) {
            res[i] = nums[i]
            res[i + nums.length] = nums[i]
        }

        return res
    }
}
