using System;
using System.Linq;


namespace AoC.PassportRules
{
    public class EyeColourPassportRule: IPassportRule
    {
        readonly  string[] _validValues = new [] {"amb", "blu", "brn", "gry", "grn", "hzl", "oth"};


        public string FieldName => "ecl";

        public bool IsValidValue(string fieldValue) => _validValues.Contains(fieldValue);
    }
}
