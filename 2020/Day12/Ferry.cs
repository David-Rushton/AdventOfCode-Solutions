using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;


namespace AoC
{
    public class Ferry: IFerry
    {
        public FerryLocation PlotCourse(IEnumerable<NavigationInstruction> Instructions)
        {
            var location = new FerryLocation(Orintation.East, 0 , 0);

            foreach(var instruction in Instructions)
            {
                Debug.Assert(instruction.Direction != NavigationDirection.Unknown, $"Unsupported direction: {instruction}");
                Debug.Assert(instruction.Value >= 0, $"Negative instruction values are not supported: {instruction}");

                Verbose.WriteLine($"Course change: {instruction}");

                location = UpdateLocation(location, instruction);
                Verbose.WriteLine($"New location: {location}\n");
            }


            return location;
        }


        private FerryLocation UpdateLocation(FerryLocation location, NavigationInstruction instruction)
        {
            var orintation = instruction.Direction switch
            {
                NavigationDirection.Left    => CalculateNewOrintation(location.Orintation, instruction.Direction, instruction.Value),
                NavigationDirection.Right   => CalculateNewOrintation(location.Orintation, instruction.Direction, instruction.Value),
                _                           => location.Orintation
            };

            (int NorthSouth, int EastWest) offset = instruction.Direction switch
            {
                NavigationDirection.Left    => (0, 0),
                NavigationDirection.Right   => (0, 0),
                _                           => CalculateTravelOffset(location, instruction)
            };


            return new FerryLocation(
                orintation,
                location.NorthSouth + offset.NorthSouth,
                location.EastWest + offset.EastWest
            );
        }

        private Orintation CalculateNewOrintation(Orintation orintation, NavigationDirection direction, int value)
        {
            var newOrintation = (int)orintation;

            newOrintation += direction == NavigationDirection.Right ? value : (value * -1);

            if(newOrintation < 0)
                newOrintation += 360;

            if(newOrintation >= 360)
                newOrintation -= 360;


            var validOrintations = new [] {0, 90, 180, 270};
            Debug.Assert(validOrintations.Contains(newOrintation), $"Invalid orintation: {newOrintation}");


            return (Orintation)newOrintation;
        }

        private (int NorthSouth, int EastWest) CalculateTravelOffset(FerryLocation location, NavigationInstruction instruction) =>
            instruction.Direction switch
            {
                NavigationDirection.North   => (instruction.Value, 0),
                NavigationDirection.East    => (0, instruction.Value),
                NavigationDirection.South   => (instruction.Value * -1, 0),
                NavigationDirection.West    => (0, instruction.Value * -1),
                NavigationDirection.Forward =>
                    location.Orintation switch
                    {
                        Orintation.North   => (instruction.Value, 0),
                        Orintation.East    => (0, instruction.Value),
                        Orintation.South   => (instruction.Value * -1, 0),
                        Orintation.West    => (0, instruction.Value * -1),
                        _                       => (0, 0)
                    },
                _                           => (0, 0),
            }
        ;
    }
}
