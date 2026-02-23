import 'package:dio/dio.dart';
import 'package:frontend/domain/entity/exchange_rate.dart';

import '../../core/api_client.dart';
import '../models/exchange_rate.dart';

abstract class ExchangeRateDataSource{
  Future<ExchangeRateEntity> getLatestExchangeRate();
  Future<String> addExchangeRate(ExchangeRateEntity exchangeRate);

}

class ExchangeRateDataSourceImpl implements ExchangeRateDataSource{
  final ApiClient apiClient;

  ExchangeRateDataSourceImpl({required this.apiClient});

  @override
  Future<ExchangeRateEntity> getLatestExchangeRate() async {
    try{
      final response = await apiClient.dio.get('/exchange-rate/latest');
      if (response.statusCode != 200) {
        throw Exception('Failed to fetch exchange rate: ${response.statusCode}');
      }

      return ExchangeRateModel.fromJSON(response.data);
    } on DioException catch (e) {
      throw Exception('Connection Failed: ${e.message}');}
    catch(e){
      throw Exception('Error fetching data: $e');


    }

  }

  @override
  Future<String> addExchangeRate(ExchangeRateEntity exchangeRate) async {
    try{


      var data = ExchangeRateModel.fromEntity(exchangeRate).toJson();

      final response = await apiClient.dio.post('/exchange-rate', data: data);

      if (response.statusCode != 201) {
        throw Exception('Failed to add exchange rate: ${response.statusCode}');
      }

      return "Exchange rate added successfully";
    }on DioException catch (e) {
      throw Exception('Connection Failed: ${e.message}');}
    catch (e){
      throw Exception("Error adding exchange rate: $e");
    }


  }


}