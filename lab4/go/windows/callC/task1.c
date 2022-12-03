#include <stdbool.h>
#include "task1.h"

double main(int selectedMaterial, int selectedSklop,
	bool sillChecked, double height, double width) {

	double sum = 0;
	if (selectedMaterial == 0 && selectedSklop == 0) {
    		sum *= 2.5;
    	}
    	else if (selectedMaterial == 0 && selectedSklop == 1) {
    		sum *= 3.0;
    	}
    	else if (selectedMaterial == 1 && selectedSklop == 0) {
    		sum *= 0.5;
    	}
    	else if (selectedMaterial == 1 && selectedSklop == 1) {
    		sum *= 1.0;
    	}
    	else if (selectedMaterial == 1.5 && selectedSklop == 0) {
    		sum *= 1.5;
    	}
    	else if (selectedMaterial == 2 && selectedSklop == 1) {
    		sum *= 2.0;
    	}
	return sum;
}