<?php

namespace App\Admin\Models;

use Illuminate\Database\Eloquent\Model;
use App\Observers\ModelObserver;


class PowCoin extends Model
{
    protected $table = 'pow_coin';

    public function algorithm() {
        return $this->belongsTo(Algorithm::class, "algorithm_id");
    }

    public static function boot() {
        parent::boot();
        $class = get_called_class();
        $class::observe(new ModelObserver());
    }
}
