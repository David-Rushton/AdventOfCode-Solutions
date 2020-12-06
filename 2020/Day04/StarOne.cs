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
            // Passport fields:
            //   byr (Birth Year)
            //   iyr (Issue Year)
            //   eyr (Expiration Year)
            //   hgt (Height)
            //   hcl (Hair Color)
            //   ecl (Eye Color)
            //   pid (Passport ID)
            //   [cid (Country ID)]

            var validPassports = reader.GetPassports().Count(p => p.HasAllRequiredFields());
            Console.WriteLine($"\nValid passports: {validPassports}");
        }
    }
}
