using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.IO;
using System.Linq;


namespace AoC
{
    public class Lexer
    {
        public IEnumerable<TileToken> GetTokens(string inputPath)
        {
            string buffer = string.Empty;

            foreach(var line in File.ReadLines(inputPath))
            {
                var directions = new List<Direction>();

                foreach(var character in line)
                    switch (character)
                    {
                        case 's':
                        case 'n':
                            buffer = character.ToString();
                            break;

                        case 'e':
                        case 'w':
                            directions.Add(GetDirection(character));
                            buffer = string.Empty;
                            break;

                        default:
                            throw new Exception($"Input not recognised: {character}");
                    }

                yield return new TileToken(Colour.White, GetLocation(directions), directions);
            }


            Direction GetDirection(char direction) =>
                ($"{buffer}{direction.ToString()}") switch
                {
                    "se" => Direction.SouthEast,
                    "sw" => Direction.SouthWest,
                    "ne" => Direction.NorthEast,
                    "nw" => Direction.NorthWest,
                    "e"  => Direction.East,
                    "w"  => Direction.West,
                    _    => throw new Exception($"Direction not supported: {buffer}{direction.ToString()}")
                }
            ;
        }


        private Location GetLocation(List<Direction> directions)
        {
            var northing = 0.0;
            var easting = 0.0;

            foreach(var direction in directions)
                switch (direction)
                {
                    case Direction.East:
                        easting++;
                        break;

                    case Direction.NorthEast:
                        northing++;
                        easting += .5;
                        break;

                    case Direction.NorthWest:
                        northing++;
                        easting -= .5;
                        break;

                    case Direction.West:
                        easting--;
                        break;

                    case Direction.SouthEast:
                        northing--;
                        easting += .5;
                        break;

                    case Direction.SouthWest:
                        northing--;
                        easting -= .5;
                        break;

                    default:
                        throw new Exception($"Direction not supported: {direction}");
                }


            return new Location(northing, easting);
        }
    }
}
