using System;
using System.Text.RegularExpressions;


namespace AoC.PassportRules
{
    public class HairColourPassportRule: IPassportRule
    {
        const string _hairColourPattern = @"^#([0-9]|[a-f]){6}$";

        readonly Regex _regEx = new Regex(_hairColourPattern);


        public string FieldName => "hcl";

        public bool IsValidValue(string fieldValue) => _regEx.IsMatch(fieldValue.Trim());
    }
}
