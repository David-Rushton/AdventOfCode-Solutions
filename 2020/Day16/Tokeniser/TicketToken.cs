namespace AoC.Tokeniser
{
    public record TicketToken(
        int TicketId,
        TicketOwner Owner,
        int[] Fields
    );
}
