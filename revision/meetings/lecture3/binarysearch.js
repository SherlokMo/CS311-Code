function binarySearch(arr, k) {

    let low = 0, high = arr.length - 1
    while(low <= high) {
        let mid = Math.floor((low+high) / 2)
        console.log("low bound is:" +low +" "+ "high bound is:" + high)
        if(arr[mid] == k) return mid
        else if(arr[mid] > k) { // [ .... , 20] new bound , 71, 80]
            high = mid - 1
        }
        else{
            low = mid + 1
        }
    }

    return -1
}

arr = [1, 2, 3, 4, 5, 6, 7, 9, 10, 11]

console.log("Index: " + binarySearch(arr, 5))