using System;
using System.Linq;


namespace  AoC
{
    public class ChineseRemainderTheorem
    {
        public long Solve(long[] numbers, long[] remainders)
        {
            long seed = 1;
            long prod = numbers.Aggregate(seed, (x, y) => x * y);
            long runningTotal = 0;

            for(long i = 0; i < numbers.Length; i ++)
            {
                long pp = numbers[i] * prod;
                long inv = ModInverse(numbers[i], pp);
                long total = remainders[i] * pp * inv;

                runningTotal += total;
            }


            return runningTotal % prod;
        }


        private long ModInverse(long i, long mod)
        {
            return Power(i, mod - 2, mod);


            long Power(long x, long y, long m)
            {
                if(y == 0)
                    return 1;

                long p = Power(x, y / 2, m) % m;
                p = (p * p) % m;

                if(y % 2 == 0)
                    return p;
                else
                    return (x * p) % m;
            }
        }
    }
}
