import React from 'react';
import { View, ScrollView, StyleSheet } from 'react-native';

export default class ListsScreen extends React.Component {
  static navigationOptions = {
    title: 'Lists',
  };

  render() {
    return (
      <View>
        <ScrollView style={styles.container}>
        
        </ScrollView>
      </View>
    );
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    paddingTop: 30,
    backgroundColor: '#fff',
  },
});
