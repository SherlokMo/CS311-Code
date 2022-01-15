#include <vector>
#include <iostream>

using namespace std;

vector<int> mergeSort(vector<int> array);
vector<int> mergeSortedArrays(vector<int> leftHalf, vector<int> rightHalf);

int main() {
    vector<int> arr = {20, 50, 670, 10, 60, 90, 5, 1, 25, 70};
    // we can reduce memory usage by sending arr by reference.
    vector<int> sortedArr = mergeSort(arr);
    for(int i = 0; i < sortedArr.size(); i++) {
        cout << sortedArr[i] << " ";
    }
}

vector<int> mergeSort(vector<int> array) {
    if(array.size() <= 1) return array;

    int mid = array.size() / 2;
    vector<int> rightHalf(array.begin() + mid, array.end());
    vector<int> leftHalf(array.begin(), array.begin() + mid);

    return mergeSortedArrays(mergeSort(leftHalf), mergeSort(rightHalf)); 
}

vector<int> mergeSortedArrays(vector<int> leftHalf, vector<int> rightHalf) {
    vector<int> sortedArray(leftHalf.size() + rightHalf.size(), 0); // Initalizie a vector with full size and fill it with zeros
   
    int i = 0;
    int j = 0;
    int k = 0;

    while(i < leftHalf.size() && j < rightHalf.size()) {
        if(leftHalf[i] <= rightHalf[j]) 
            sortedArray[k++] = leftHalf[i++];
        else 
            sortedArray[k++] = rightHalf[j++];
    }

    while(i < leftHalf.size()) {
        sortedArray[k++] = leftHalf[i++];
    }
    while(j < rightHalf.size()) {
        sortedArray[k++] = rightHalf[j++];
    }

    return sortedArray;
}