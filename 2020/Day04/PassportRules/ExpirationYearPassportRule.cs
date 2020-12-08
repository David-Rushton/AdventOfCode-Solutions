using System;


namespace AoC.PassportRules
{
    public class ExporationYearPassportRule: IPassportRule
    {
        public string FieldName => "eyr";

        public bool IsValidValue(string fieldValue) =>
            int.Parse(fieldValue) is >= 2020 and <= 2030
        ;
    }
}
