using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;


namespace AoC
{
    public enum FerryOrintation
    {
        North = 0,
        East = 90,
        South = 180,
        West = 270
    }

    public record FerryLocation(
        FerryOrintation Orintation,
        int NorthSouth,
        int EastWest
    );


    public class Ferry
    {
        public FerryLocation PlotCourse(IEnumerable<NavigationInstruction> Instructions)
        {
            var location = new FerryLocation(FerryOrintation.East, 0 , 0);

            foreach(var instruction in Instructions)
            {
                Debug.Assert(instruction.Direction != NavigationDirection.Unknown, $"Unsupported direction: {instruction}");
                Debug.Assert(instruction.Value >= 0, $"Negative instruction values are not supported: {instruction}");

                location = UpdateLocation(location, instruction);
            }


            return location;
        }


        private FerryLocation UpdateLocation(FerryLocation location, NavigationInstruction instruction)
        {
            var orintation = instruction.Direction switch
            {
                NavigationDirection.North   => FerryOrintation.North,
                NavigationDirection.East    => FerryOrintation.East,
                NavigationDirection.South   => FerryOrintation.South,
                NavigationDirection.West    => FerryOrintation.West,
                NavigationDirection.Left    => CalculateNewOrintation(location.Orintation, instruction.Direction, instruction.Value),
                NavigationDirection.Right   => CalculateNewOrintation(location.Orintation, instruction.Direction, instruction.Value),
                _                           => location.Orintation
            };

            (int NorthSouth, int EastWest) offset = instruction.Direction switch
            {
                NavigationDirection.Left    => (0, 0),
                NavigationDirection.Right   => (0, 0),
                _                           => CalculateTravelOffset(orintation, instruction.Value)
            };


            return new FerryLocation(
                orintation,
                location.NorthSouth + offset.NorthSouth,
                location.EastWest + offset.EastWest
            );
        }

        private FerryOrintation CalculateNewOrintation(FerryOrintation orintation, NavigationDirection direction, int value)
        {
            var newOrintation = (int)orintation;

            if(direction == NavigationDirection.Right)
                value = value *1;


            if(newOrintation < 0)
                newOrintation += 360;

            if(newOrintation > 360)
                newOrintation -= 360;


            var validOrintations = new [] {0, 90, 180, 270};
            Debug.Assert(validOrintations.Contains(newOrintation), $"Invalid orintation: {newOrintation}");


            return (FerryOrintation)newOrintation;
        }

        private (int NorthSouth, int EastWest) CalculateTravelOffset(FerryOrintation orintation, int value) =>
            orintation switch
            {
                FerryOrintation.North   => (value, 0),
                FerryOrintation.East    => (0, value),
                FerryOrintation.South   => (value * -1, 0),
                FerryOrintation.West    => (0, value * -1),
                _                       => (0, 0),
            }
        ;
    }
}
