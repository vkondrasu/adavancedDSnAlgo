class HashRing:
    def __init__(self, k):
        self.head = None
        self.k = k
        self.min = 0
        self.max = 2**k - 1

    def legalRange(self, hashValue):
        return self.min <= hashValue <= self.max

    def distance(self, resourceHashValue, nodeHashValue):
        if resourceHashValue == nodeHashValue :
            return 0
        if nodeHashValue > resourceHashValue:
            return nodeHashValue-resourceHashValue
        else:
            return (2 ** self.k) + (nodeHashValue-resourceHashValue)

    def lookupNode(self, hashValue):
        if self.legalRange(hashValue):
            temp = self.head

            if temp is None:
                return None
            else:
                while self.distance(temp.hashValue, hashValue) > self.distance(temp.next.hashValue, hashValue):
                    temp = temp.next
                if temp.hashValue == hashValue:
                    return temp
                return temp.next

    def moveResources(self, dest, orig, deleteTrue):
        delete_list = []
        for i,j in orig.resource.items():
            if self.distance(i,dest.hashValue) < self.distance(i, orig.hashValue) or deleteTrue:
                dest.resource[i] = j
                delete_list.append(i)
            for i in delete_list:
                del orig.resources[i]

class Node:
    def __init__(self, hashValue):
        self.hashValue = hashValue
        self.resource = {}
        self.next = None
        self.previous = None

