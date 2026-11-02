from __future__ import annotations
from typing import List, Optional

import sys


class MinHeap:
    """
    Abstracting the node data as Int but could be Any, given a custom comparison
    function to guide ourselves on how to compare the nodes.
    """
    nodes: List[int]

    def __init__(self, nodes: List[int] = []):
        self.nodes = []
        for node in nodes:
            self.add(node)

    def __len__(self) -> int:
        return len(self.nodes)

    def __get_left_child_index(self, parent_index: int) -> int:
        return 2 * parent_index + 1

    def __get_right_child_index(self, parent_index: int) -> int:
        return 2 * parent_index + 2

    def __get_parent_index(self, child_index: int) -> int:
        return (child_index - 1) // 2

    def __has_left_child(self, parent_index: int) -> bool:
        return self.__get_left_child_index(parent_index) < len(self.nodes)

    def __has_right_child(self, parent_index: int) -> bool:
        return self.__get_right_child_index(parent_index) < len(self.nodes)

    def __has_parent(self, index: int) -> bool:
        return self.__get_parent_index(index) >= 0

    def __left_child(self, index: int) -> Optional[int]:
        if not self.__has_left_child(index):
            # Virtual -inf
            return -sys.maxsize
        return self.nodes[self.__get_left_child_index(index)]

    def __right_child(self, index: int) -> Optional[int]:
        if not self.__has_right_child(index):
            # Virtual -inf
            return -sys.maxsize
        return self.nodes[self.__get_right_child_index(index)]

    def __parent(self, index: int) -> Optional[int]:
        if not self.__has_parent(index):
            return None
        return self.nodes[self.__get_parent_index(index)]

    def __swap(self, first_idx: int, second_idx: int):
        if first_idx >= len(self.nodes) or second_idx >= len(self.nodes):
            print(f'first ({first_idx}) or second ({second_idx}) are invalid')
            return
        tmp = self.nodes[first_idx]
        self.nodes[first_idx] = self.nodes[second_idx]
        self.nodes[second_idx] = tmp

    def EstaVacio(self) -> bool:
        return not self.nodes

    def VerTope(self) -> Optional[int]:
        if not self.is_empty():
            return None
        return self.nodes[0]

    def heapify_down(self, index: Optional[int] = 0):
        """Move root to proper position in heap"""
        if index >= len(self.nodes) or not self.__has_left_child(index):
            return self.nodes
        # Check if greater than left or right
        smaller_child_idx = self.__get_left_child_index(index)
        if self.__has_right_child(index) and \
                self.__right_child(index) < self.__left_child(index):
            smaller_child_idx = self.__get_right_child_index(index)

        if self.nodes[index] < self.nodes[smaller_child_idx]:
            # Lower than both, do nothing
            return self.nodes

        if self.nodes[index] > self.nodes[smaller_child_idx]:
            # Swap with smaller child
            self.__swap(index, smaller_child_idx)
            return self.heapify_down(smaller_child_idx)

    def heapify_up(self, child: Optional[int] = None):
        """Move last added element to correct position in heap"""
        if not child:
            # Start with last
            child = len(self.nodes) - 1
        # Check if smaller than parent
        parent_idx = self.__get_parent_index(child)
        if self.__parent(child) and self.nodes[child] < self.__parent(child):
            # Swap child <> parent
            self.__swap(child, parent_idx)
            # Recurse on parent index now (should have child value)
            self.heapify_up(child=parent_idx)

        # If its not smaller, leave as it is
        return self.nodes

    def Desencolar(self) -> Optional[int]:
        """
        Polls lowest element from the heap (usually root). To avoid memory shift
        of first-element removal, we copy the last element to the first
        position, shrink the size by 1 and heapify down.
        """
        if self.EstaVacio():
            print('Empty, not polling')
            return None

        # Remove first and insert the last one added as the root
        removed_node = self.nodes[0]
        self.nodes[0] = self.nodes[len(self.nodes) - 1]
        # Could've also used self.nodes = self.nodes[:-1]
        del self.nodes[-1]

        self.heapify_down()
        return removed_node

    def Encolar(self, item: int):
        self.nodes.append(item)
        self.heapify_up()

    def heapify_children(self, index: int):
        """
        Heapify children of a node to obey left < right.
        """
        if not self.__has_left_child(index):
            return
        if not self.__has_right_child(index):
            return
        if self.__right_child(index) < self.__left_child(index):
            self.__swap(
                self.__get_right_child_index(index),
                self.__get_left_child_index(index))
        self.heapify_children(self.__get_left_child_index(index))
        self.heapify_children(self.__get_right_child_index(index))


def heapsort_in_place(unsorted_input: List[int]) -> List[int]:
    """ Heapsort in-place: heap.nodes is unsorted_input == True """
    heap = MinHeap()
    heap.nodes = unsorted_input
    size = len(unsorted_input)
    print(f'identity check: {heap.nodes is unsorted_input}')
    for idx in range(size // 2, -1, -1):
        heap.heapify_down(idx)
    heap.heapify_children(0)
    return heap.nodes


def heapsort_aux(unsorted_input: List[int]) -> List[int]:
    """ Heapsort using O(n) space """
    heap = MinHeap(unsorted_input)
    sorted_input = []
    for _ in range(len(unsorted_input)):
        sorted_input.append(heap.poll())
    print(f'identity check: {sorted_input is unsorted_input}')
    return sorted_input