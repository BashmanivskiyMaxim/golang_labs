#include <stdbool.h>
#include "task2.h"

double main(int days, int Country, int Season, bool guide, bool luxury) {
	double sum = 0;

	if (Country == 0 && Season == 0) {
		sum = days * 100.0;
	}
	else if (Country == 0 && Season == 1) {
		sum = days * 150.0;
	}
	else if (Country == 1 && Season == 0) {
		sum = days * 160.0;
	}
	else if (Country == 1 && Season == 1) {
		sum = days * 200.0;
	}
	else if (Country == 2 && Season == 0) {
		sum = days * 120.0;
	}
	else if (Country == 2 && Season == 1) {
		sum = days * 180.0;
	}

	if (guide == true) {
    		sum = sum + days * 50.0;
    	}

	if (luxury == true) {
		sum = sum + sum * 0.2;
	}

	return sum;
}