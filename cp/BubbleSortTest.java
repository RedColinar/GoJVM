package cp;

public class BubbleSortTest {
	public static void main(String[] args) {
		//
		int[] arr = {41,15,22,19,95,11,11,18,15};
		bubbleSort(arr);
		printArray(arr);
	}

	private static void bubbleSort(int[] arr){
		boolean swapped = true;
		int j = 0;
		int tmp;
		while (swapped){
			swapped = false;
			j++;
			for(int i = 0;i < arr.length - j; i++){
				if (arr[i] > arr[i + 1]){
					tmp = arr[i];
					arr[i] = arr[i + 1];
					arr[i + 1] = tmp;
					swapped = true;
				}
			}
		}
	}

	private static void printArray (int[] arr){
		System.out.print("[");
		for(int i : arr){
			System.out.print(i+" ");
		}
		System.out.print("]");
	}
} 