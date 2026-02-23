
import 'package:dio/dio.dart';
import 'package:frontend/core/api_client.dart';

import '../../domain/entity/address.dart';
import '../models/address.dart';

abstract class AddressDataSource{
  Future<AddressEntity> getAddress(String cep);
}

class AddressDataSourceImpl implements AddressDataSource {
  final ApiClient apiClient;

  AddressDataSourceImpl({required this.apiClient});

  @override
  Future<AddressEntity> getAddress(String cep) async {
    try {
      final response = await apiClient.dio.get('/zipcode/$cep');

      if (response.statusCode != 200) {
        throw Exception('Failed to fetch address: ${response.statusCode}');
      }

      return AddressModel.fromJson(response.data);
    } on DioException catch (e) {
      throw Exception('Connection Failed: ${e.message}');
    } catch (e) {
      throw Exception('Error fetching data: $e');
    }
  }
}