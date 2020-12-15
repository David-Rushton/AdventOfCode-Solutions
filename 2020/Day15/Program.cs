using System;
using System.Collections.Generic;
using System.Linq;


// T1: The 1st number spoken is a starting number, 0.
// T2: The 2nd number spoken is a starting number, 3.
// T3: The 3rd number spoken is a starting number, 6.
// T4: Last number spoken, 6. First time the number had been spoken, the 4th number spoken is 0.
// T5: Last number spoken, 0. Spoken before, the next number is the difference (T4) (T1).
// T6: The last number spoken, 3 had also been spoken before, most recently on turns 5 and 2. So, the 6th number spoken is 5 - 2, 3.
// T7: Since 3 was just spoken twice in a row, and the last two turns are 1 turn apart, the 7th number spoken is 1.
// T8: Since 1 is new, the 8th number spoken is 0.
// T9: 0 was last spoken on turns 8 and 4, so the 9th number spoken is the difference between them, 4.
// T10: 4 is new, so the 10th number spoken is 0.

var startingNumbers = new Queue<int>(GetStartingNumbers());
var numberHistory = new Dictionary<int, Number>();
Number lastNumber = new Number(0, 0, 0, false);


for(var turn = 1; turn <= 30_000_000; turn++)
{
    var number = GetNumber(turn);
    updateNumberHistory(turn, number);


    if(turn >= 30_000_000 || turn % 1_000_000 == 0)
        Console.WriteLine($"#{turn.ToString("00")}: {numberHistory[number.N]}");

    lastNumber = numberHistory[number.N];
}



IEnumerable<int> GetStartingNumbers()
{
    foreach(var arg in args)
        if(int.TryParse(arg, out var number))
            yield return number;
}

Number GetNumber(int turn)
{
    // Always return starting number first
    if(startingNumbers.Count > 0)
        return new Number(startingNumbers.Dequeue(), turn, -1, true);

    // If last turn introduced a new number then we return 0 this time
    if(numberHistory[lastNumber.N].NewNumber)
        return numberHistory[0];

    // Last turn was not a new number, return the difference between last two appearances
    var difference = lastNumber.LastSpoken - lastNumber.LastSpokenButOne;

    if(numberHistory.ContainsKey(difference))
        return numberHistory[difference];
    else
        return new Number(difference, turn, -1, true);
}

void updateNumberHistory(int turn, Number number)
{
    if(numberHistory.ContainsKey(number.N))
        numberHistory[number.N] = new Number (number.N, turn, number.LastSpoken, false);
    else
        numberHistory.Add(number.N, new Number (number.N, turn, -1, true));
}

public record Number(
    int N,
    int LastSpoken,
    int LastSpokenButOne,
    bool NewNumber
);
