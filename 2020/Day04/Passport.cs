using System;
using System.Collections.Generic;
using System.Linq;
using AoC.PassportRules;


namespace AoC
{
    public class Passport
    {
        readonly Dictionary<string, IPassportRule> _passportRules = new Dictionary<string, IPassportRule>();


        public Passport() => Fields = new Dictionary<string, string>();


        public Dictionary<string, string> Fields { get; set; }


        public void AddRule(IPassportRule passportRule) => _passportRules.Add(passportRule.FieldName, passportRule);

        public bool HasAllRequiredFields()
        {
            // Passport fields.
            // All but cid are required.
            //   byr (Birth Year)
            //   iyr (Issue Year)
            //   eyr (Expiration Year)
            //   hgt (Height)
            //   hcl (Hair Color)
            //   ecl (Eye Color)
            //   pid (Passport ID)
            //   [cid (Country ID)]
            var requiredFields = new []
            {
                "byr", "iyr", "eyr",
                 "hgt", "hcl", "ecl", "pid"
            }.ToList();

            var requiredFieldsFound = Fields.Count(f => requiredFields.Contains(f.Key));

            return (requiredFieldsFound == 7);
        }

        public override string ToString()
        {
            var result = "";
            foreach(var field in Fields)
            {
                result += $"{field.Key}: {field.Value} ";
            }

            return $"{HasAllRequiredFields()}: {result}";
        }
    }
}
