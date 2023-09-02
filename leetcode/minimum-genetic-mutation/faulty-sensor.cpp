// 1826. Faulty Sensor
// https://leetcode.com/problems/faulty-sensor/

#include <bits/stdc++.h>

using namespace std;


class Solution {
public:
    int badSensor(vector<int>& sensor1, vector<int>& sensor2) {
        int n = sensor1.size();
        int i = 0;
        while (i < n && sensor1[i] == sensor2[i])
            ++i;
        
        if (i == n)
            return -1;

        vector<int> v1, v2;

        v1 = {sensor1.begin(), sensor1.end() - 1};
        v2 = {sensor2.begin(), sensor2.end() - 1};
        if (v1 == v2)
            return -1; 
        
        bool can_be_one = false, can_be_two = false;

        v1 = {sensor1.begin() + i, sensor1.end() - 1};
        v2 = {sensor2.begin() + i + 1, sensor2.end()};
        if (v1 == v2)
            can_be_one = true;
        
        v1 = {sensor1.begin() + i + 1, sensor1.end()};
        v2 = {sensor2.begin() + i, sensor2.end() - 1};
        if (v1 == v2)
            can_be_two = true;

        if (can_be_one && can_be_two)
            return -1;
        else if (can_be_one)
            return 1;
        else if (can_be_two)
            return 2;
        else
            return -1;
    }
};
