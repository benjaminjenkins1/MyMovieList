import React from 'react';
import Colors from '../constants/Colors';
import { View, ScrollView, StyleSheet } from 'react-native';

export default class SettingsScreen extends React.Component {
  static navigationOptions = {
    title: 'Search',
    headerStyle: {
      backgroundColor: Colors.tabBar,
      borderBottomWidth: 0
    },
    headerTintColor: Colors.tintColor
  };

  render() {
    return (
      <View style={styles.container}>

      </View>
    )
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    paddingTop: 30,
    backgroundColor: Colors.contentBackground,
  },
});