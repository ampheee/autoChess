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

void enterData(int firstMatrix[][10], int secondMatrix[][10], int matrixSize) {
	cout << endl << "Enter elements of matrix 1:" << endl;
	for(int i = 0; i < matrixSize; ++i) {
		for(int j = 0; j < matrixSize; ++j) {
			cout << "Enter elements MatrixA"<< i + 1 << j + 1 << ": ";
			cin >> firstMatrix[i][j];
		}
	}
	cout << endl << "Enter elements of matrix 2:" << endl;
	for(int i = 0; i < matrixSize; ++i) {
		for(int j = 0; j < matrixSize; ++j) {
			cout << "Enter elements MatrixB" << i + 1 << j + 1 << ": ";
			cin >> secondMatrix[i][j];
		}
	}
}

void multiplyMatrices(int firstMatrix[][10], int secondMatrix[][10], int mult[][10], int matrixSize) {
	for(int i = 0; i < matrixSize; ++i) {
		for(int j = 0; j < matrixSize; ++j) {
			mult[i][j] = 0;
		}
	}
	for(int i = 0; i < matrixSize; ++i) {
		for(int j = 0; j < matrixSize; ++j) {
			for(int k=0; k < matrixSize; ++k) {
				mult[i][j] += firstMatrix[i][k] * secondMatrix[k][j];
			}
		}
	}
}

void show(int mult[][10], int matrixSize) {
	cout << "Multiplied Matrix:" << endl;
	for(int i = 0; i < matrixSize; ++i) {
		for(int j = 0; j < matrixSize; ++j) {
			cout << mult[i][j] << " ";
			if(j == matrixSize - 1)
				cout << endl;
		}
	}
}