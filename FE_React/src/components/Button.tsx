import { StyleSheet, Text, TouchableOpacity, View } from "react-native";
import React from "react";
import { useNavigation } from "@react-navigation/native";
import AsyncStorage from "@react-native-async-storage/async-storage";

export default function MyButton({
  text,
  background,
  textColor,
  onPress,
  navigateTo,
}: {
  text: string;
  background: string;
  textColor: string;
  onPress?: () => void;
  navigateTo?: string | undefined;
}) {
  const navigation = useNavigation();

  const handlePress = async () => {
    if (onPress) {
      await onPress();
    }

    if (navigateTo) {
      navigation.navigate(navigateTo as never);
    }
  };

  return (
    <TouchableOpacity
      onPress={handlePress}
      style={{
        backgroundColor: background,
        paddingVertical: 15,
        width: 300,
        borderRadius: 10,
        marginTop: 20,
        alignItems: "center",
        alignSelf: "auto",
        elevation: 5,
      }}
    >
      <Text style={{ color: textColor, fontWeight: "bold", fontSize: 17 }}>
        {text}
      </Text>
    </TouchableOpacity>
  );
}

const styles = StyleSheet.create({});
