#include <stdio.h>
#include <stdlib.h>
#include <sys/types.h>
#include <unistd.h>
#include <pthread.h>

int getRandomIterCount() {
  return rand() % 100 + 300;
}

void busyWork(int iterCount) {
  int i, j, mul;
  FILE *devNull;

  for (i = 0; i < iterCount; i++) {
    for (j = 0; j < iterCount; j++) {
      devNull = fopen("/dev/null", "w");
			mul = i * j;
      fprintf(devNull, "%d", mul);
      fclose(devNull);
    }
  }
}

void say(char word[], int index) {
  int iterCount;

  iterCount = getRandomIterCount();

  busyWork(iterCount);

  printf("Hello, %s! from %d say; Iter Count: %d\n", word, index+1, iterCount);
}

// void *sayThreaded(void *vargp) {
//   say("world", 41);
// }

int main() {
  int i;

	for (i = 0; i < 4; i++) {
    say("world", i);
  }

  // ------------- Process --------------
  // int isGrandParent, isParent;
  // isGrandParent = fork();

  // say("world", 41);

  // if(isGrandParent == 0) {
  //   isParent = fork();

  //   say("world", 41);

  //   if(isParent != 0) {
  //     wait(0);
  //   }
  // } else {
  //   wait(0);
  // }

  // ------------- Thread -------------

  // pthread_t thread_ids[4];
	// for (i = 0; i < 4; i++) {
  //   pthread_create(&thread_ids[i], NULL, sayThreaded, NULL);
  // }

	// for (i = 0; i < 4; i++) {
  //   pthread_join(thread_ids[i], NULL);
  // }

  return 0;
}
