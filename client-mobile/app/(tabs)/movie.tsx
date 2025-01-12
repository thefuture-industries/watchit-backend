import { Colors } from '@/constants/Colors';
import { View, Text, StyleSheet } from 'react-native';

export default function MovieScreen() {
  return (
    <View>
      <Text className="text-3xl" style={styles.text}>
        Movie Screen
      </Text>
    </View>
  );
}

const styles = StyleSheet.create({
  text: {
    color: Colors.app.text,
  },
});
