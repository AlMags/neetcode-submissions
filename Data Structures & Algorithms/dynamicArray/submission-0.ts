class DynamicArray {
    private array: number[];
    private capacity: number;
    private size: number;
    
    /**
     * @constructor
     * @param {number} capacity
     */
    constructor(capacity: number) {
        this.array = new Array(capacity).fill(0);
        this.capacity = capacity;
        this.size = 0;
    }

    /**
     * @param {number} i
     * @returns {number}
     */
    get(i: number): number {
        return this.array[i];
    }

    /**
     * @param {number} i
     * @param {number} n
     * @returns {void}
     */
    set(i: number, n: number): void {
        this.array[i] = n;

    }

    /**
     * @param {number} n
     * @returns {void}
     */
    pushback(n: number): void {
        if (this.size === this.capacity) {
            this.resize();
        }
        this.array[this.size] = n;
        this.size++;
    }

    /**
     * @returns {number}
     */
    popback(): number {
        if (this.size != 0) {
            this.size--;
        } 
        return this.array[this.size];
    }

    /**
     * @returns {void}
     */
    resize(): void {
        this.capacity *= 2;
        let newArray = new Array(this.capacity).fill(0);

        for (let i = 0; i < this.size; i++) {
            newArray[i] = this.array[i];
        }

        this.array = newArray;
    }

    /**
     * @returns {number}
     */
    getSize(): number {
        return this.size;
    }

    /**
     * @returns {number}
     */
    getCapacity(): number {
        return this.capacity;
    }
}
