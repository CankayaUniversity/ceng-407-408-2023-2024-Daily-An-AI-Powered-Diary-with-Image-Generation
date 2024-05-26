import React from 'react';
import { useState } from 'react';
import { View, StyleSheet, Text, Image, ScrollView } from 'react-native';
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

   const { data, isLoading, isError } = useGetStatistics();

   const markedDatesFunc = data?.date.reduce((acc: any, date: any) => {
      acc[date] = { selected: true };
      return acc;
   }, {});

   return (
      <Header navigation={navigation} previous="Home" homepage={false}>
         {!isLoading &&
            <ScrollView style={{ height: '100%', width: '100%' }}>
               <View style={styles.container}>
                  <Calendar
                     key={uuidv4()}

                     markedDates={
                        markedDatesFunc
                     }

                     style={{
                        backgroundColor: Colors.main_container,
                        width: '100%',
                        paddingLeft: 0,
                        paddingRight: 0
                     }}

                     onDayPress={day => {
                        console.log(day);
                     }}

                     theme={{
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
                              color: "white"
                           },
                        },
                        'stylesheet.calendar.main': {
                           dayContainer: {
                              flex: 1,
                              alignItems: 'center'
                           },
                           week: {
                              marginVertical: "1%",
                              marginHorizontal: 0,
                              flexDirection: 'row',
                              justifyContent: 'center'
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
                        }
                     }
                     }
                     hideArrows={true}
                  />

                  <View style={styles.outerRow}>
                     <View style={styles.innerItem}>
                        <View>
                           <Text style={styles.innerText}>
                              Wows:
                           </Text>
                           <View style={{ width: "100%", height: "31%" }}></View>
                           <View style={{ flexDirection: "row", alignItems: "center" }} >
                              <Text style={styles.innerNumber}>
                                 {data?.likes || "0"}
                              </Text>
                           </View>
                        </View>
                     </View>
                     <View style={styles.innerItem}>
                        <View>
                           <Text style={styles.innerText}>
                              Dailies written:
                           </Text>
                           <View style={{ width: "100%", height: "31%" }}></View>
                           <View style={{ flexDirection: "row", alignItems: "center" }} >
                              <Text style={styles.innerNumber}>
                                 {data?.dailiesWritten || 0}
                              </Text>
                           </View>
                        </View>
                     </View>
                  </View>
                  <View style={styles.outerRow}>
                     <View style={styles.innerItem}>
                        <View>
                           <Text style={styles.innerText}>
                              Views:
                           </Text>
                           <View style={{ width: "100%", height: "31%" }}></View>
                           <View style={{ flexDirection: "row", alignItems: "center" }} >
                              <Text style={styles.innerNumber}>
                                 {data?.views || 0}
                              </Text>
                           </View>
                        </View>
                     </View>
                     <View style={styles.innerItem}>
                        <View>
                           <Text style={styles.innerText}>
                              Streak:
                           </Text>
                           <View style={{ width: "100%", height: "31%" }}></View>
                           <View style={{ flexDirection: "row", alignItems: "center" }} >
                              <Text style={styles.innerNumber}>
                                 {data?.streak || 0}
                              </Text>
                           </View>
                        </View>
                     </View>
                  </View>
                  <View style={styles.outerRow}>
                     <View style={styles.innerItem}>
                        <View>
                           <Text style={styles.innerText}>
                              Mood:
                           </Text>
                           <View style={{ width: "100%", height: "20%" }}></View>
                           <View style={{ flexDirection: "row", width: "70%", height: "56%", alignItems: "flex-end" }} >
                              <Text style={[styles.innerNumber, { fontSize: 20}]}>
                                 {data?.mood || "Uncalculated"}
                              </Text>
                           </View>
                        </View>
                     </View>
                     <View style={styles.innerItem}>
                        <View>
                           <Text style={styles.innerText}>
                              Topic:
                           </Text>
                           <View style={{ width: "100%", height: "20%" }}></View>
                           <View style={{ flexDirection: "row", width: "70%", height: "56%", alignItems: "flex-end" }} >
                              <Text style={[styles.innerNumber],{ fontSize: 20}}>
                                 {data?.topic || "Uncalculated"}
                              </Text>
                           </View>
                        </View>
                     </View>
                  </View>
               </View>
            </ScrollView>
         }
      </Header >
   );
}

const styles = StyleSheet.create({
   container: {
      alignItems: 'center',
      padding: "3%",
      paddingBottom:30,
      margin: "3%",
      borderRadius: 20,
      backgroundColor: Colors.main_container,
      opacity: 0.90,
   },
   outerRow: {
      height: '15%',
      marginTop: "6%",
      width: '92%',
      flexDirection: 'row',
      justifyContent: 'space-between',
   },
   innerItem: {
      padding: 10,
      borderRadius: 10,
      width: '45%',
      marginHorizontal: 10,
      backgroundColor: "rgba(0,0,0,0.6)",
      flexDirection: "column",
      height: '100%',
      justifyContent:"space-evenly"
   },
   innerText: {
      fontSize: 20,
      color: "rgba(255, 255, 255, 1)",
      fontWeight:'200'
   },
   innerNumber: {
      fontSize: 40,
      color: "white",
      textAlign:"left",
      fontWeight:'200'
   }

})
export default Statistics;
