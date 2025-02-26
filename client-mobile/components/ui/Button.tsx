import { View, Text, TouchableOpacity, StyleSheet } from 'react-native';
import React from 'react';
import { Colors } from '@/constants/Colors';

interface ButtonProps {
	text: string;
}

export default function Button({ text }: ButtonProps) {
	return (
		<TouchableOpacity style={styles.button}>
			<Text className="text-[#fff] text-[16px] text-center font-semibold">
				{text}
			</Text>
			<View
				style={[styles.button_shadow, { width: '100%', height: 10 }]}
			></View>
		</TouchableOpacity>
	);
}

const styles = StyleSheet.create({
	button: {
		backgroundColor: Colors.app.button_background,
		borderRadius: 7,
		position: 'relative',
		paddingTop: 13,
		paddingBottom: 14,
		marginTop: 10,
		marginBottom: 10,
	},
	button_shadow: {
		position: 'absolute',
		backgroundColor: '#000',
		bottom: -1,
		borderRadius: 2,
		opacity: 0.5,
	},
});
