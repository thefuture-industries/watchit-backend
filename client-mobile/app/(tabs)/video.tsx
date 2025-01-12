import { Colors } from '@/constants/Colors';
import { View, Text, StyleSheet } from 'react-native';

export default function VideoScreen() {
  return (
    <View>
      <Text className="text-3xl" style={styles.text}>
        Video Screen
      </Text>
    </View>
  );
}

const styles = StyleSheet.create({
  text: {
    color: Colors.app.text,
  },
});
