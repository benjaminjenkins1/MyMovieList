import React from 'react';
import { StyleSheet, Text } from 'react-native';

export default function TabBarLabel(props) {
  return(
    <Text style={[styles.tabBarLabel,  props.focused? styles.tabBarLabelActive : {}]} >{props.title}</Text>
  );
}

const styles = StyleSheet.create({
  tabBarLabel: {
    paddingBottom: 6,
    fontSize: 10,
    textAlign: 'center'
  },
  tabBarLabelActive: {
    color: 'red'
  }
});