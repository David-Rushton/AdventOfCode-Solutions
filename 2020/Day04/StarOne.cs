using System;
using System.Collections.Generic;
using System.Linq;
using System.IO;


namespace AoC
{
    public class StarOne: IStar
    {
        public void Invoke(PassportReader reader)
        {
            var result = reader.ProcessPassports();

            Console.WriteLine($"\nValid passports: {result.countOfPassportsWithAllRequiredFields}");
        }
    }
}
