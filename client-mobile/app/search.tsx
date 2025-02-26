import { Colors } from '@/constants/Colors';
import { ArrowLeft, History } from 'lucide-react-native';
import { useState } from 'react';
import {
	Text,
	View,
	TextInput,
	StyleSheet,
	TouchableOpacity,
} from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';

export default function SearchScreen() {
	const [search, setChangeSearch] = useState<string>('');

	return (
		<SafeAreaView className="h-full">
			<View className="my-4 px-4">
				<View style={{ position: 'relative' }}>
					<TextInput
						className="bg-[#444] rounded-3xl text-[#fff] pl-[40px]"
						value={search}
						onChangeText={setChangeSearch}
						placeholder="Search movie"
						placeholderTextColor="#999"
					/>
				</View>
				<TouchableOpacity
					onPress={() => {}}
					style={{
						position: 'absolute',
						left: 27,
						top: 9,
					}}
				>
					<ArrowLeft color={Colors.app.tint} size={20} />
				</TouchableOpacity>
			</View>

			<View className="my-4 px-4">
				<TouchableOpacity className="my-3 flex-row items-center">
					<View className="bg-[#222] p-3 rounded-[100px]">
						<History color={Colors.app.button_background} />
					</View>

					<Text className="text-[#fff] text-[16px] ml-5">
						Iron Man 2
					</Text>
				</TouchableOpacity>

				<TouchableOpacity className="my-3 flex-row items-center">
					<View className="bg-[#222] p-3 rounded-[100px]">
						<History color={Colors.app.button_background} />
					</View>

					<Text className="text-[#fff] text-[16px] ml-5">
						Iron Man 2
					</Text>
				</TouchableOpacity>

				<TouchableOpacity className="my-3 flex-row items-center">
					<View className="bg-[#222] p-3 rounded-[100px]">
						<History color={Colors.app.button_background} />
					</View>

					<Text className="text-[#fff] text-[16px] ml-5">
						Iron Man 2
					</Text>
				</TouchableOpacity>
			</View>

			<View className="my-4 px-4">
				<Text className="text-[#fff] text-[24px] font-semibold">
					Popular
				</Text>

				<View></View>
			</View>
		</SafeAreaView>
	);
}

const styles = StyleSheet.create({
	icon_search: {
		position: 'absolute',
		bottom: '50%',
		left: 26,
		top: 9,
	},
});
