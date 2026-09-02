class Solution {
    /**
     * @param {number[]} nums
     * @param {number} k
     * @return {number[]}
     */
    topKFrequent(nums, k) {
        const count = {};
        const frequencies = Array.from({length: nums.length + 1}, () => []);

        for(const num of nums) {
            count[num] = (count[num] || 0) + 1;
        }

        for(const num in count) {
            frequencies[count[num]].push(parseInt(num));
        }

        const result = [];

        for(let i = frequencies.length - 1; i > 0; i--) {
            for(const num of frequencies[i]) {
                result.push(num);
                if(result.length === k) {
                    return result;
                }
            }
        }
        
    }
}
