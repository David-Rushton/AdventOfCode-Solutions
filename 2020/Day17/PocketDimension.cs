using System;
using System.Collections.Generic;
using System.Linq;


namespace AoC
{
    public class PocketDimension
    {
        readonly Dictionary<string, ConwayCube> _universe;


        public PocketDimension(IEnumerable<ConwayCube> initialState) =>
            _universe = initialState.ToDictionary(k => k.Id, v => v)
        ;


        public void ExecuteBootSequence(int cycleCount)
        {
            ExpandUniverse(cycleCount);

            for(var cycle = 0; cycle < cycleCount; cycle++)
            {
                var flipToInactive = GetActiveCubesWithLessThanTwoOrMoreThanThreeActiveNeighbours();
                var flipToActive = GetInactiveCubesWithThreeActiveNeighbours();

                foreach(var cube in flipToActive.Union(flipToInactive).ToList())
                    _universe[cube.Id] = cube with { IsActive = ! cube.IsActive };
            }

            PrintUniverse();


            IEnumerable<ConwayCube> GetActiveCubesWithLessThanTwoOrMoreThanThreeActiveNeighbours() =>
                _universe.Where(c => c.Value.IsActive && CountOfActiveNeighbouringCubes(c.Value) is <2 or >3).Select(c => c.Value)
            ;

            IEnumerable<ConwayCube> GetInactiveCubesWithThreeActiveNeighbours() =>
                _universe.Where(c => c.Value.IsActive == false && CountOfActiveNeighbouringCubes(c.Value) == 3).Select(c => c.Value)
            ;
        }

        // Section of Universe that contains active cubes can grow during each cycle.
        // Maximum growth is +1 to each boundary.
        // Ex universe of 3x3x1 becomes 5x5x3, then 7x7x5 etc.
        // Pre-sizing the universe simplifies processing and visualisations.
        private void ExpandUniverse(int cycleCount)
        {
            var newX = GetRange( _universe.Select(c => c.Value.Position.X).Max() );
            var newY = GetRange( _universe.Select(c => c.Value.Position.Y).Max() );
            var newZ = GetRange( _universe.Select(c => c.Value.Position.Z).Max() );

            foreach(var x in newX)
                foreach(var y in newY)
                    foreach(var z in newZ)
                        createCube(x, y, z);


            IEnumerable<int> GetRange(int dimensionMax) =>
                Enumerable.Range(0 - cycleCount, dimensionMax + (cycleCount * 2) +  1)
            ;

            void createCube(int x, int y, int z)
            {
                var cube = new ConwayCube(new ConwayCubePosition(x, y, z), false);
                if( ! _universe.ContainsKey(cube.Id) )
                    _universe.Add(cube.Id, cube);
            }
        }

        private int CountOfActiveNeighbouringCubes(ConwayCube cube) =>
            GetNeighbouringCubes(cube).Where(c => c.IsActive).Count()
        ;

        private IEnumerable<ConwayCube> GetNeighbouringCubes(ConwayCube cube)
        {
            foreach(var neighbourPosition in GetNeighbouringCubePositions(cube.Position))
                yield return GetCubeOrReturnDefault(neighbourPosition);


            // Cubes outside the universe are always inactive
            ConwayCube GetCubeOrReturnDefault(ConwayCubePosition position) =>
                _universe.ContainsKey(position.Id)
                ? _universe[position.Id]
                : new ConwayCube(position, false)
            ;
        }

        private IEnumerable<ConwayCubePosition> GetNeighbouringCubePositions(ConwayCubePosition cubePosition)
        {
            for(int xOffset = -1; xOffset < 2; xOffset++)
                for(int yOffset = -1; yOffset < 2; yOffset++)
                    for(int zOffset = -1; zOffset < 2; zOffset++)
                        if( ! (xOffset == 0 && yOffset == 0 && zOffset == 0) )
                            yield return getConwayCubePositionOffset(xOffset, yOffset, zOffset);


            ConwayCubePosition getConwayCubePositionOffset(int x, int y, int z) =>
                new ConwayCubePosition(cubePosition.X + x, cubePosition.Y + y, cubePosition.Z + z)
            ;
        }

        private void PrintUniverse()
        {
            var zSeries = _universe.Select(kvp => kvp.Value.Position.Z).Distinct().OrderBy(z => z);

            foreach(var z in zSeries)
            {
                var sequence = _universe
                    .Where(c => c.Value.Position.Z == z)
                    .OrderBy(c => c.Value.Position.Y)
                    .ThenBy(c => c.Value.Position.X)
                    .Select
                    (
                        c =>
                        new { Y = c.Value.Position.Y, X = c.Value.Position.X, State = c.Value.IsActive ? '#' : '.'  }
                    )
                ;
                var lastY = sequence.Min(c => c.Y);

                Console.WriteLine($"\nZ: {z}");

                foreach(var cell in sequence)
                {
                    if(cell.Y != lastY)
                    {
                        Console.WriteLine();
                        lastY = cell.Y;
                    }
                    Console.Write(cell.State);
                }

                Console.WriteLine();
            }


            var countOfActiveCubes = _universe.Where(c => c.Value.IsActive).Count();
            Console.WriteLine($"\nActive cubes in universe {countOfActiveCubes}\n");
        }
    }
}
