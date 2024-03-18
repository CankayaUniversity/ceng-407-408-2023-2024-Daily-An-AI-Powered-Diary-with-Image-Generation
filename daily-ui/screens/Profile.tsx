import { StatusBar } from 'expo-status-bar';
import { StyleSheet, Text, View,Alert,TouchableOpacity, ImageBackground } from 'react-native';
import Header from '../components/Header';

const Profile = ({navigation}:{navigation:any}) => {
   return (
      <Header navigation={navigation} previous="Home" homepage={true}>
            <View style={styles.container}>
               <View style={styles.container2}>
                  <View style={styles.rows}>
                     <View style={{width:'40%', paddingTop:12,borderRadius:12, alignItems: 'center',backgroundColor:'rgba(0,0,0,0.4)',aspectRatio: 1/1}}>
                        <View style={styles.avatarPlaceholder}/>
                     </View>
                     <View style={{width:'60%', paddingTop:12,borderRadius:12,backgroundColor:'rgba(0,0,0,0.4)'}}>
                        <Text style={styles.name}>daily dev</Text>
                        <Text style={styles.username}>@daily_dev</Text>
                     </View>
                  </View>
                  <View style={styles.rows}>
                     <View style={styles.box2}>
                        <Text style={styles.username}>test</Text>
                     </View>  
                  </View>
                  <View style={styles.rows}>
                     <View style={styles.box1}>
                        <Text style={styles.username}>test</Text>
                     </View>
                     <View style={styles.box1}>
                        <Text style={styles.username}>test</Text>
                     </View>  
                  </View>
                  <View style={styles.rows}>
                     <View style={styles.box1}>
                        <Text style={styles.username}>test</Text>
                     </View>
                     <View style={styles.box1}>
                        <Text style={styles.username}>test</Text>
                     </View>  
                  </View>
                  
               </View>
            </View>
      </Header>
   );
}

const styles = StyleSheet.create({
   container: {
      height: '90%',
      padding:12
   },
   container2: {
      height: '100%',
      borderRadius:12,
      padding:12,
      backgroundColor:'rgba(0,0,0,0.4)'
   },
   rows: {
      justifyContent: 'space-around',
      height: '25%',
      marginBottom: 12,
      flex: 1,
      flexDirection: 'row',
      borderRadius:12,
   },
   box1: {
      paddingStart:12,
      paddingEnd:12,
      paddingTop:12, 
      borderTopStartRadius:12,
      borderTopEndRadius:12,
      width:'50%', 
      borderRadius:12,
      backgroundColor:'rgba(0,0,0,0.4)' 
   },
   box2: {
      paddingStart:12,
      paddingEnd:12,
      paddingTop:12, 
      borderTopStartRadius:12,
      borderTopEndRadius:12,
      width:'100%', 
      borderRadius:12,
      backgroundColor:'rgba(0,0,0,0.4)'
   },
   name: {
      textAlign:'right',
      marginEnd:12.5,
      fontSize:36,
      fontWeight:'500',
      color:'white',
      opacity: 1.0
   },
   username: {
      textAlign:'right',
      marginEnd:12.5,
      fontSize:18,
      fontWeight:'200',
      color:'white',
      opacity: 0.5
   },
   avatar: {
      width: 128,
      height: 128,
      borderRadius: 50,
   },
   avatarPlaceholder: {
      width: 128,
      height: 128,
      borderRadius: 100,
      backgroundColor: '#FFFFFF',
      opacity: 0.5,
      alignSelf: 'center',
      justifyContent: 'center',
   }
   
});

export default Profile;
