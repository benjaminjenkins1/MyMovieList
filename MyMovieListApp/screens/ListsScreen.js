import React from 'react';
import { View, ScrollView, StyleSheet } from 'react-native';
import Colors from '../constants/Colors';

export default class ListsScreen extends React.Component {
  static navigationOptions = {
    title: 'Lists',
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
    );
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    paddingTop: 30,
    backgroundColor: Colors.contentBackground,
  },
});
