using System;


namespace AoC.PassportRules
{
    public class IssueYearPassportRule: IPassportRule
    {
        public string FieldName => "iyr";

        public bool IsValidValue(string fieldValue) =>
            int.Parse(fieldValue) is >= 2010 and <= 2020
        ;
    }
}
