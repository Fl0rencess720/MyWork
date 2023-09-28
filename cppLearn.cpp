#include <iostream>
using namespace std;

int main()
{
   float a = 32.89;
   int b = 133345;
   char x[] = {"4444422222"};
   char y[] = {"abcdefj"};
   printf("%5.1f,%09d\n", a, b);
   printf("%-9s,%-10s\n", x, y);
   printf("%09s是%4.4s\n", "aceTaffy", "lostLove");
   // char ch="\100";
   string ch;
   ch = "怎么会这样";
   cout << ch << endl;
   int d, e;
   scanf("%d,%d", &d, &e);
   printf("%d就是我%d", d, e);

   return 0;
}