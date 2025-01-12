import { ReactNode } from 'react';
import { View } from 'react-native';
import {
  GestureHandlerRootView,
  TapGestureHandler,
  TapGestureHandlerGestureEvent,
} from 'react-native-gesture-handler';
import Animated, {
  measure,
  runOnJS,
  useAnimatedGestureHandler,
  useAnimatedStyle,
  useSharedValue,
  withTiming,
} from 'react-native-reanimated';

interface RippleAnimationProps {
  children: ReactNode;
  ref: any;
  onTap?: () => void;
}

export default function RippleAnimation({
  children,
  ref,
  onTap,
}: RippleAnimationProps) {
  // **
  // Animation
  const centerX = useSharedValue(0);
  const centerY = useSharedValue(0);
  const width = useSharedValue(0);
  const height = useSharedValue(0);
  const scale = useSharedValue(0);
  const opacityMovie = useSharedValue(0);

  const tapGestureEvent =
    useAnimatedGestureHandler<TapGestureHandlerGestureEvent>({
      onStart: (tapEvent) => {
        const layout = measure(ref);

        if (layout) {
          width.value = layout.width;
          height.value = layout.height;
        }

        centerX.value = tapEvent.x;
        centerY.value = tapEvent.y;

        opacityMovie.value = 1;
        scale.value = 0;
        scale.value = withTiming(1, { duration: 300 });
      },
      onActive: () => {
        if (onTap) runOnJS(onTap)();
      },
      onFinish: () => {
        opacityMovie.value = withTiming(0);
      },
    });

  const rStyle = useAnimatedStyle(() => {
    const circleRadius = Math.sqrt(width.value ** 2 + height.value ** 2);

    const translateX = centerX.value - circleRadius;
    const translateY = centerY.value - circleRadius;

    return {
      width: circleRadius * 2,
      height: circleRadius * 2,
      opacity: opacityMovie.value,
      top: 0,
      left: 0,
      borderRadius: circleRadius,
      backgroundColor: 'rgba(0, 0, 0, 0.5)',
      position: 'absolute',
      transform: [
        { translateX },
        { translateY },
        {
          scale: scale.value,
        },
      ],
    };
  });

  return (
    <View ref={ref}>
      <GestureHandlerRootView>
        <TapGestureHandler onGestureEvent={tapGestureEvent}>
          <Animated.View style={{ overflow: 'hidden' }}>
            {children}
            <Animated.View style={rStyle} />
          </Animated.View>
        </TapGestureHandler>
      </GestureHandlerRootView>
    </View>
  );
}
