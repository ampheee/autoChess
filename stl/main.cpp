#include "multiplyMatrices.h"


int main()
{   int matricesSize = 0;
	int firstMatrix[10][10], secondMatrix[10][10], mult[10][10];
	cout << "Enter matrixs size:\n";
	cin >> matricesSize;
        enterData(firstMatrix, secondMatrix, matricesSize);
        multiplyMatrices(firstMatrix, secondMatrix, mult, matricesSize);
        show(mult, matricesSize);
	return 0;
}
