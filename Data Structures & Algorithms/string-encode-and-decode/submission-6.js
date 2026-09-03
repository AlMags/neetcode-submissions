class Solution {
    /**
     * @param {string[]} strs
     * @returns {string}
     */
    encode(str) {
        if (str.length === 0) {
            return String(JSON.stringify(str));
        }
        if (str.length === 1 && str[0] === '') {
            return ""
        }
        const mergedArray = str.join('#/#')
        let encodedArray = new Array(mergedArray.length).fill('')
        
        for(let i = 0; i < encodedArray.length; i++) {
            encodedArray[i] = mergedArray[i].charCodeAt(0);
        }

        return encodedArray.join(','); 
    }

    /**
     * @param {string} str
     * @returns {string[]}
     */
    decode(str) {
        if (str === "[]") {
            return []
        }
        if (str === '') {
            return [""]
        }
        const strToArr = str.split(',')
        let decodedString = new Array(strToArr.length).fill('');
        let arr = [];

        for(let i = 0; i < strToArr.length; i++) {
            decodedString[i] = String.fromCharCode(strToArr[i]);
        }

        arr = decodedString.join('').split('#/#')

        return arr;
    }
}
