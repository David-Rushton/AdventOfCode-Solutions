using System;


namespace AoC.PassportRules
{
    public interface IPassportRule
    {
        string FieldName { get; }

        public bool IsValidValue(string fieldValue);
    }
}
