using System;


namespace AoC.PassportRules
{
    public class BirthYearPassportRule: IPassportRule
    {
        public string FieldName => "byr";

        public bool IsValidValue(string fieldValue) => int.Parse(fieldValue) is >= 1920 and <= 2002;
    }
}
