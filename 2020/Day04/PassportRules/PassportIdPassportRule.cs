using System;
using System.Text.RegularExpressions;


namespace AoC.PassportRules
{
    public class PassportIdPassportRule: IPassportRule
    {
        const string _passportIdPattern = @"^[0-9]{9}$";

        readonly Regex _regEx = new Regex(_passportIdPattern);


        public string FieldName => "pid";

        public bool IsValidValue(string fieldValue) => _regEx.IsMatch(fieldValue);
    }
}
