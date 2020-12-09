using System;


namespace AoC
{
    public record Token(
        string Container,
        string Contained,
        int ContainedCount
    );
}
