import sys
from enum import Enum
from dataclasses import dataclass
from typing import Iterator

FaceValuesMap = {
    'A': '14',
    'K': '13',
    'Q': '12',
    'T': '10',
    '9': '09',
    '8': '08',
    '7': '07',
    '6': '06',
    '5': '05',
    '4': '04',
    '3': '03',
    '2': '02',
    'J': '01'
}

class HandType(Enum):
    FIVE_OF_A_KIND = 7
    FOUR_OF_A_KIND = 6
    FULL_HOUSE = 5
    THREE_OF_A_KIND = 4
    TWO_PAIR = 3
    ONE_PAIR = 2
    HIGH_CARD = 1


@dataclass
class Hand:
    cards: str
    bid: int
    hand_type: HandType
    rank: int
    score: int


def main(is_test_mode: bool) -> None:
    print('Camel Cards')
    print()

    hands = list(get_input(is_test_mode))
    for hand in hands:
        score_hand(hand)

    sorted_hands = sorted(hands, key=lambda x: x.score, reverse=False)
    rank = 1
    total_winning = 0
    for hand in sorted_hands:
        hand.rank = rank
        rank += 1
        print(f' {hand.cards} {hand.score} {hand.rank} {hand.bid * hand.rank}')
        total_winning = total_winning + (hand.bid * hand.rank)

    print()
    print(f'Winning: {total_winning}')


def score_hand(hand: Hand):
    score = str(hand.hand_type.value)
    for card in hand.cards:
        score += FaceValuesMap[card]
    hand.score = int(score)

def get_input(get_test: bool) -> Iterator[Hand]:
    test_path = './input.test.txt'
    prod_path = './input.txt'
    path = test_path if get_test else prod_path

    for line in open(path, 'rt').read().splitlines():
        content = line.split(' ')
        yield Hand(
            content[0],
            int(content[1]),
            get_hand_type(content[0]),
            0,
            0)


def get_hand_type(hand: str) -> HandType:
    cards: dict[str, int] = {}
    jokers_found = 0
    for card in hand:
        if card in cards:
            cards[card] += 1
        else:
            cards[card] = 1
        if card == 'J':
            jokers_found += 1

    contains_3_of_a_kind = False
    pairs_found = 0
    for key, value in sorted(cards.items()):
        if value == 5:
            return joker_upgrade(HandType.FIVE_OF_A_KIND, jokers_found)
        elif value == 4:
            return joker_upgrade(HandType.FOUR_OF_A_KIND, jokers_found)
        elif value == 3:
            contains_3_of_a_kind = True
        elif value == 2:
            pairs_found += 1

    if contains_3_of_a_kind:
        if pairs_found == 1:
            return joker_upgrade(HandType.FULL_HOUSE, jokers_found)
        else:
            return joker_upgrade(HandType.THREE_OF_A_KIND, jokers_found)

    if pairs_found > 0:
        if pairs_found == 2:
            return joker_upgrade(HandType.TWO_PAIR, jokers_found)
        else:
            return joker_upgrade(HandType.ONE_PAIR, jokers_found)

    return joker_upgrade(HandType.HIGH_CARD, jokers_found)


def joker_upgrade(hand_type: HandType, number_of_jokers: int) -> HandType:
    if number_of_jokers == 0:
        return hand_type
    # FIVE_OF_A_KIND = 7    joker  no improvement
    if hand_type == HandType.FIVE_OF_A_KIND:
        return hand_type
    # FOUR_OF_A_KIND = 6    joker  if 1 j then 5oak, if 4 js then 5oak
    if hand_type == HandType.FOUR_OF_A_KIND:
        if number_of_jokers == 1:
            return HandType.FIVE_OF_A_KIND
        elif number_of_jokers == 4:
            return HandType.FIVE_OF_A_KIND
    # FULL_HOUSE = 5        joker  if 2 j then 5oak, if 3 js then 5oak
    if hand_type == HandType.FULL_HOUSE:
        if number_of_jokers == 2:
            return HandType.FIVE_OF_A_KIND
        elif number_of_jokers == 3:
            return HandType.FIVE_OF_A_KIND
    # THREE_OF_A_KIND = 4   joker  if 1 j then 4oak, if 3 js then 4oak
    if hand_type == HandType.THREE_OF_A_KIND:
        if number_of_jokers == 1:
            return HandType.FOUR_OF_A_KIND
        elif number_of_jokers == 3:
            return HandType.FOUR_OF_A_KIND
    # TWO_PAIR = 3          joker  if 1 j then full hse, if 2 j then 4oak
    if hand_type == HandType.TWO_PAIR:
        if number_of_jokers == 1:
            return HandType.FULL_HOUSE
        elif number_of_jokers == 2:
            return HandType.FOUR_OF_A_KIND
    # ONE_PAIR = 2          joker  if 2 j then 3oak, if 1 j then 3oak
    if hand_type == HandType.ONE_PAIR:
        if number_of_jokers == 2:
            return HandType.THREE_OF_A_KIND
        elif number_of_jokers == 1:
            return HandType.THREE_OF_A_KIND
    # HIGH_CARD = 1         joker  if 1 j then one pair
    if hand_type == HandType.HIGH_CARD:
        if number_of_jokers == 1:
            return HandType.ONE_PAIR

    raise('I did not think we could get here :(')


if __name__ == '__main__':
    is_test_mode = sys.argv.count('test') > 0
    main(is_test_mode)
