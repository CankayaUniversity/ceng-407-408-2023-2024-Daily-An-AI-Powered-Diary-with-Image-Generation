import React from 'react';
import { useState } from 'react';
import { TouchableWithoutFeedback, Modal, View, StyleSheet, Text, ScrollView, Pressable, TouchableOpacity } from 'react-native';
import Header from '../components/Header';
import { Calendar } from 'react-native-calendars';
import { Colors } from '../libs/colors.tsx';
import { v4 as uuidv4 } from 'uuid';
import { useGetStatistics } from '../libs';
import { StatisticsResponse } from '../libs/types';

const calendarSelectedColor = '#AF5C5C66';

const Statistics = ({ navigation }: { navigation: any }) => {
    const [timeframe, setTimeframe] = useState("month");
    const [statistics, setStatistics] = useState<StatisticsResponse | null>(null);
    const [error, setError] = useState<string | null>(null);
    const [isModalVisible, setModalVisible] = useState(false);

    const { data, isLoading, isError } = useGetStatistics();

    const markedDatesFunc = data?.date?.reduce((acc: any, date: any) => {
        acc[date] = { selected: true };
        return acc;
    }, {});

    return (
        <Header navigation={navigation} previous="Home" homepage={false}>
            {(!isLoading && data?.date != null) && (
                <View style={styles.scrollView}>
                    <View style={styles.container}>
                        <Calendar
                            key={uuidv4()}
                            markedDates={markedDatesFunc}
                            style={styles.calendar}
                            onDayPress={(day) => {
                                console.log(day);
                            }}
                            theme={calendarTheme}
                            hideArrows={true}
                        />

                        <View style={styles.outerRow}>
                            <View style={styles.innerItem}>
                                <Text style={styles.innerText}>Wows:</Text>
                                <Text style={styles.innerNumber}>{data?.likes || "0"}</Text>
                            </View>
                            <View style={styles.innerItem}>
                                <Text style={styles.innerText}>Dailies written:</Text>
                                <Text style={styles.innerNumber}>{data?.dailiesWritten || 0}</Text>
                            </View>
                        </View>
                        <View style={styles.outerRow}>
                            <View style={styles.innerItem}>
                                <Text style={styles.innerText}>Views:</Text>
                                <Text style={styles.innerNumber}>{data?.views || 0}</Text>
                            </View>
                            <View style={styles.innerItem}>
                                <Text style={styles.innerText}>Streak:</Text>
                                <Text style={styles.innerNumber}>{data?.streak || 0}</Text>
                            </View>
                        </View>
                        <View style={[styles.outerRow,{height:'15%'}]}>
                            <View style={styles.innerItem}>
                                <Text style={styles.innerText}>Mood:</Text>
                                <Text style={[styles.innerNumber, { fontSize: 20 }]}>{data?.mood.toUpperCase() || "Uncalculated"}</Text>
                            </View>
                            <Pressable onPress={() => setModalVisible(!isModalVisible)} style={styles.innerItem}>
                                <Text style={styles.innerText}>Topic:</Text>
                                <Text style={[styles.innerNumber, { fontSize: 20 }]}>{data?.topics[0].toUpperCase() || "Uncalculated"}</Text>
                            </Pressable>
                        </View>
                    </View>
                </View>
            )}
            <Modal
                visible={isModalVisible}
                transparent={true}
                animationType="slide"
                onRequestClose={() => setModalVisible(false)}
                style={styles.modal}
            >
                <View style={styles.modalOverlay}>
                    <View style={styles.modalContentWrapper}>
                        <ScrollView style={styles.modalContent} contentContainerStyle={styles.modalContentContainer}>
                            {data?.topics?.map((el, index) => (
                                <TouchableOpacity key={index} activeOpacity={1} onPress={() => setModalVisible(!isModalVisible)} style={styles.modalItem}>
                                    <Text style={styles.modalItemText}>{el.toUpperCase()}</Text>
                                </TouchableOpacity>
                            ))}
                        </ScrollView>
                    </View>
                </View>
            </Modal>
        </Header>
    );
};

const styles = StyleSheet.create({
    scrollView: {
        flex: 1,
        width: '100%',
    },
    container: {
        alignItems: 'center',
        padding:15,
        margin: "2%",
        borderRadius: 20,
        backgroundColor: Colors.main_container,
        opacity: 0.90,
    },
    calendar: {
        backgroundColor: Colors.main_container,
        width: '100%',
        borderRadius: 20,
        paddingLeft: 0,
        paddingRight: 0,
    },
    outerRow: {
        marginTop: "6%",
        width: '92%',
        height: '15%',
        flexDirection: 'row',
        justifyContent: 'space-between',
    },
    innerItem: {
        padding: 10,
        borderRadius: 10,
        width: '45%',
        height:'100%',
        marginHorizontal: 10,
        backgroundColor: "rgba(0,0,0,0.6)",
        flexDirection: "column",
        justifyContent: 'space-between',
    },
    innerText: {
        fontSize: 20,
        color: "rgba(255, 255, 255, 1)",
        fontWeight: '200',
    },
    innerNumber: {
        fontSize: 32,
        color: "white",
        textAlign: "left",
        fontWeight: '200',
    },
    modal: {
        height: '100%',
    },
    modalOverlay: {
        flex: 1,
        justifyContent: 'center',
        alignItems: 'center',
        backgroundColor: 'rgba(0, 0, 0, 0.5)',
    },
    modalContentWrapper: {
        height: "60%",
        width: "90%",
    },
    modalContent: {
        width: '100%',
        height: '100%',
        backgroundColor: Colors.main_container,
        borderRadius: 10,
        borderWidth: 0.5,
        borderColor: 'gray',
    },
    modalContentContainer: {
        padding: 5,
        gap: 5,
    },
    modalItem: {
        backgroundColor: Colors.main_container,
        height: 40,
        justifyContent: 'center',
    },
    modalItemText: {
        fontSize: 24,
        color: 'white',
        textAlign: 'center',
        fontWeight: '200',
    },
});

const calendarTheme = {
    textDayFontSize: 20,
    textMonthFontSize: 20,
    selectedDayTextColor: 'white',
    calendarBackground: 'transparent',
    monthTextColor: Colors.background,
    dayTextColor: 'white',
    'stylesheet.calendar.header': {
        fontSize: 20,
        width: '100%',
        alignItems: 'center',
        marginTop: 5,
        flexDirection: 'row',
        justifyContent: 'center',
        dayHeader: {
            marginBottom: "1%",
            width: "12%",
            textAlign: 'center',
            fontSize: 16,
            color: "white",
        },
    },
    'stylesheet.calendar.main': {
        dayContainer: {
            flex: 1,
            alignItems: 'center',
        },
        week: {
            marginVertical: "1%",
            marginHorizontal: 0,
            flexDirection: 'row',
            justifyContent: 'center',
        },
    },
    'stylesheet.day.basic': {
        selected: {
            color: calendarSelectedColor,
            backgroundColor: calendarSelectedColor,
            width: "95%",
            height: 36,
            borderRadius: 10,
        },
    },
};

export default Statistics;
